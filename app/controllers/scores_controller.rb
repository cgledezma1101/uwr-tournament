class ScoresController < ApplicationController
	load_and_authorize_resource
	before_action :authenticate_user!

	# POST /scores
	#
	# Method called when a player scores a goal on a game
	#
	# @param [Integer] game_id Identifier of the game where the score was made
	# @param [Integer] player_id Identifier of the player that scored
	def create
		@game = @score.game
		@player = @score.player

		@score.save

		@player_goals = @score.game.goals_for(@player)
		@team_goals = @score.game.scores.joins(:player).where(player: { team_id: @player.team_id}).count
	end

	private

	# Sanitizes parameters for creating a new score
	def create_params
		params.require(:score).permit(:game_id, :player_id)
	end
end
