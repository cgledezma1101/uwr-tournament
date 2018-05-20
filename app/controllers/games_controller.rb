class GamesController < ApplicationController
	load_and_authorize_resource
	before_action :authenticate_user!

	# GET /games/add_chronometer
	#
	# Returns javascript appropriate for adding a chronometer to the game that performs the request
	def add_chronometer
		respond_to do |format|
			format.js { render }
		end
	end

	# POST /games
	#
	# Based on the parameters received, creates a new game and redirects the user
	# to the screen that visualizes it.
	def create
		@game.status = Game::STATUS_READY
		if @game.save
			redirect_params = { notice: t('game.create_success') }
		else
			alert_message = "#{t('game.create_fail')} #{stringify_errors(@game.errors, 'game')}"
			redirect_params = { alert: alert_message }
		end

		redirect_to stage_path(@game.stage), redirect_params
	end

	# DELETE /games/:id
	#
	# Destroys the specified instance of a game
	#
	# @param [Integer] id Identifier of the game to be destroyed
	def destroy
		@game.destroy
		redirect_to stage_path(@game.stage), notice: t('game.destroy_success')
	end

	# GET /games/:id/end
	#
	# Request a modal with the information necessary to end a game
	#
	# @param [Integer] id Identifier of the game to end
	def end
		render 'games/_end_game.html.haml', layout: false
	end

	# GET /games/external_scoreboard
	#
	# Renders a form that contains a layout for rendering a standalone scoreboard for the game
	def external_scoreboard
	end

	# POST /games/:id/finalize
	#
	# Marks the specified game as ended, so that no further changes can be done on it
	#
	# @param [Integer] id Identifier of the game to mark as ended
	def finalize
		blue_tournament_team = @game.stage.tournament.tournament_teams.where(team: @game.blue_team).first
		white_trounament_team = @game.stage.tournament.tournament_teams.where(team: @game.white_team).first

		blue_password_valid = blue_tournament_team.valid_password?(params[:game][:blue_password])
		white_password_valid = white_trounament_team.valid_password?(params[:game][:white_password])

		if !blue_password_valid && !white_password_valid
			redirect_to game_path(@game), alert: t('game.errors.finalize.both_mismatch')
			return
		elsif !blue_password_valid
			redirect_to game_path(@game), alert: t('game.errors.finalize.team_mismatch', team_name: @game.blue_team.name)
			return
		elsif !white_password_valid
			redirect_to game_path(@game), alert: t('game.errors.finalize.team_mismatch', team_name: @game.white_team.name)
			return
		end

		@game.status = Game::STATUS_ENDED
		white_goals = @game.white_goals
		blue_goals = @game.blue_goals

		if white_goals > blue_goals
			winning_color = PlayerGame::WHITE_TEAM
		elsif blue_goals > white_goals
			winning_color = PlayerGame::BLUE_TEAM
		else
			winning_color = Game::TIED_GAME
		end

		@game.update(status: Game::STATUS_ENDED, winning_color: winning_color)
		redirect_to game_path(@game), notice: t('game.finalize_success')
	end


	# GET /games/new?stage_id=:stage_id
	#
	# Renders a form that would allow the creation of a game in a stage
	#
	# @param [Integer] stage_id Identifier of the stage to which this game will be added
	def new
		stage = Stage.find(params[:stage_id])
		authorize! :update, stage

		@game.stage = stage
		@teams = stage.tournament.teams.order(:name)

		render 'games/_new', layout: false
	end

	# DELETE /games/:id/remove_score
	#
	# Allows removing a goal from a player in this game
	#
	# @param [Integer] id Identifier if the game where the change is being made
	# @param [Integer] game.player_id Identifier of the player that will loose a score
	def remove_score
		@player = Player.find(params[:game][:player_id])
		score_to_remove = @game.scores.where{ player_id == my{@player.id} }.first
		if !score_to_remove.nil?
			score_to_remove.destroy
		end

		@player_goals = @game.goals_for(@player)
		@team_goals = @game.scores.joins{ player }.where{ player.team_id == my{@player.team_id} }.count
	end

	# GET /games/:id
	#
	# Displays a game and allows basic mannipulation over it
	#
	# @param [Integer] id The identifier of the game to be displayed
	def show
		@score = Score.new
	end

	# POST /game/:id/start
	#
	# Marks the given game as started and associates all the currently active players on each team to their respective colors
	#
	# @param [Integer] id Identifier of the game that will be started
	def start
		rolled_back = false
		ActiveRecord::Base.transaction do
			for blue_player in @game.blue_team.active_players do
				relation = PlayerGame.new(game: @game, player: blue_player, team_color: PlayerGame::BLUE_TEAM)
				if !relation.save
					rolled_back = true
					raise ActiveRecord::Rollback
				end
			end

			for white_player in @game.white_team.active_players do
				relation = PlayerGame.new(game: @game, player: white_player, team_color: PlayerGame::WHITE_TEAM)
				if !relation.save
					rolled_back = true
					raise ActiveRecord::Rollback
				end
			end

			@game.status = Game::STATUS_STARTED

			if !@game.save
				rolled_back = true
				raise ActiveRecord::Rollback
			end
		end

		if rolled_back
			redirect_params = { alert: t('game.game_start_error') }
		else
			redirect_params = { notice: t('game.game_started') }
		end

		redirect_to game_path(@game), redirect_params
	end

	private
		# This defines the attributes that are permitted when creating a new game
		def create_params
			params.require(:game).permit(:blue_team_id, :white_team_id, :stage_id, :starts_at, :blue_team_calculation, :white_team_calculation)
		end
end
