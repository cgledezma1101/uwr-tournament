class TournamentInvitationsController < ApplicationController
	before_action :authenticate_user!

	# POST /tournament_invitations
	#
	# Allows the creation of an invitation of a club to a tournament
	#
	# @param [Integer] club_id Identifier of the club to be invited
	# @param [Integer] tournament_id Identifier of the tournament the club is being invited to
	def create
		invitation = TournamentInvitation.new(create_params)
		authorize! :update, invitation.tournament

		if invitation.save
			redirect_params = { notice: t('tournament_invitation.create_success', club_name: invitation.club.name, tournament_name: invitation.tournament.name) }
		else
			redirect_params = { alert: t('tournament_invitation.create_failure') }
		end

		redirect_to tournament_path(invitation.tournament), redirect_params
	end

	# DELETE /tournament_invitations/:id
	#
	# Destroys the instance of the tournament invitation with the specified id
	#
	# @param [Integer] id Identifier of the tournament invitation to be destroyed
	def destroy
		invitation = TournamentInvitation.find(params[:id])
		authorize! :update, invitation.tournament

		invitation.destroy

		redirect_to tournament_path(invitation.tournament), notice: t('tournament_invitation.destroy_success', club_name: invitation.club.name, tournament_name: invitation.tournament.name)
	end

	# GET /tournament_invitations?tournament_id=:tournament_id
	#
	# Renders a form that allows an administrator to invite a club to a tournament
	#
	# @param [Integer] tournament_id Identifier of the tournament for which the invitation will be created
	def new
		tournament = Tournament.find(params[:tournament_id])
		authorize! :update, tournament

		@tournament_invitation = TournamentInvitation.new(tournament: tournament)

		tournament_clubs_ids = tournament.clubs.select{ |club| club.id }
		@eligible_clubs = Club.where{ id << my{tournament_clubs_ids} }

		render 'tournament_invitations/_new', layout: false
	end

	private

	# Sanitized parameters to create a new tournament invitation
	def create_params
		params.require(:tournament_invitation).permit(:club_id, :tournament_id)
	end
end
