class TournamentTeamsController < ApplicationController
	before_action :authenticate_user!

	# POST /tournament_teams
	#
	# Allows the creation of the relation between a team and a tournament
	#
	# @param [Integer] team_id Identifier of the team to be sent to the tournament
	# @param [Integer] tournament_id Identifier of the tournament the team will be sent to
	def create
		tournament_team = TournamentTeam.new(create_params)
		authorize! :update, tournament_team.team

		if tournament_team.save
			redirect_params = { notice: I18n.t('tournament_team.save_success', team_name: tournament_team.team.name, tournament_name: tournament_team.tournament.name) }
		else
			redirect_params = { alert: I18n.t('tournament_team.save_fail', team_name: tournament_team.team.name, tournament_name: tournament_team.tournament.name) }
		end

		redirect_to club_path(tournament_team.team.club), redirect_params
	end

	# GET /tournament_teams/new?team_id=:team_id
	#
	# Renders a form that allows sending a team to a tournament
	#
	# @param [Integer] team_id Identifier of the team that the user wishes to send to a tournament
	def new
		team = Team.find(params[:team_id])
		authorize! :update, team

		invited_tournaments = Tournament.joins(:clubs).where(clubs: { id: team.club_id }).to_a
		tournaments_going = Tournament.joins(:teams).where(teams: { id: team.id }).to_a
		@available_tournaments = invited_tournaments - tournaments_going
		@tournament_team = TournamentTeam.new(team: team)

		render 'tournament_teams/_new', layout: false
	end

	private

	# Returns the sanitized parameters required to create a relationship between a team and a tournament
	def create_params
		params.require(:tournament_team).permit(:tournament_id, :team_id, :password, :password_confirmation)
	end
end
