class TeamsController < ApplicationController
	before_action :authenticate_user!
	load_and_authorize_resource

	# GET /teams/:id/players
	#
	# For the specified team, searches it's players and renders a form associated
	# to them. This action answers esclusively to JS calls.
	#
	# @param [Integer] id Identifier of the team whos players are desired
	def players
		@players = @team.players
		@game = Game.new
	end

	# GET /teams/new?club_id=:club_id
	#
	# Renders the form that will allow the creation of a new team for a club
	#
	# @param [Integer] club_id Identifier of the club to which the team will be added
	def new
		@club = Club.find(params[:club_id])
		authorize! :update, @club

		render 'teams/_new', layout: false
	end

	# POST /teams
	#
	# Allows for the creation of a new team, associating it to a club
	def create
		if(@team.save)
			redirect_to team_path(@team)
		else
			redirect_to club_path(@team.club), alert: t('team.standard_save_error')
		end
	end

	# GET /teams/:id
	#
	# Displays the general information of a particular team
	def show
		@club = @team.club
	end

	# GET /teams/:id/confirm_destroy
	#
	# Presents a dialog that allows the deletion of a team
	#
	# @param [Integer] id Identifier of the team to be destroyed
	def confirm_destroy
		render 'teams/_confirm_destroy', layout: false
	end

	# DELETE /teams/:id
	#
	# Removes the specified team from the database
	def destroy
		@team.destroy
		redirect_to club_path(@team.club), notice: t('team.destroy_successful')
	end

	# GET /teams/:id/edit
	#
	# Presents a form that allows the modification of the teams information
	#
	# @param [Integer] id Identifier of the team to be edited
	def edit
		@club = @team.club
		render 'teams/_new', layout: false
	end

	# POST /teams/:id
	#
	# Allows the update of the teams information
	#
	# @param [Integer] id Identifier of the team to be updated
	# @param [String] team[name] New name to give to the team
	def update
		if(@team.update(update_params))
			redirect_params = { notice: t('team.update_success') }
		else
			redirect_params = { alert: "#{t('team.update_failure')}: #{stringify_errors(@team.errors, 'team')}" }
		end

		redirect_to team_path(@team), redirect_params
	end

	private

	# Parameter sanitizer for the create method
	def create_params
		params.require(:team).permit(:name, :club_id)
	end

	# Parameter sanitizer for the update method
	def update_params
		params.require(:team).permit(:name)
	end
end
