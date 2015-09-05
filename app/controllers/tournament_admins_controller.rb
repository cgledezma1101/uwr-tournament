class TournamentAdminsController < ApplicationController
  before_action :authenticate_user!

  # GET /tournament_admins/new?tournament_id=:tournament_id
  #
  # Presents an interface to add an administrator to the specified tournament
  #
  # @param [Integer] tournament_id Identifier of the tournament for which the administrator will be added
  def new
    tournament = Tournament.find(params[:tournament_id])
    authorize! :update, tournament

    @tournament_admin = TournamentAdmin.new(tournament: tournament)
    tournament_admin_ids = tournament.admins.to_a.select{ |admin| admin.id }
    @eligible_users = User.where{ id << my{tournament_admin_ids} }
    render 'tournament_admins/_new', layout: false
  end
end
