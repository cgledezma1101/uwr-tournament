class TournamentInvitationsController < ApplicationController
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
end
