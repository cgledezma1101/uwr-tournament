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

  # POST /tournament_admins
  #
  # Performs the creation of a tournaments administrator
  #
  # @param [Integer] tournament_id Identifier of the tournament to be associated
  # @param [Integer] user_id Identifier of the user  to be associated
  def create
    tournament = Tournament.find(params[:tournament_admin][:tournament_id])
    authorize! :update, tournament

    relationship = TournamentAdmin.new(create_params)
    if relationship.save
      redirect_params = { notice: t('tournament_admin.create_success', user_name: relationship.user.name, tournament_name: tournament.name) }
    else
      redirect_params = { alert: t('tournament_admin.create_fail') }
    end

    redirect_to tournament_path(tournament), redirect_params
  end

  private

  def create_params
    params.require(:tournament_admin).permit(:tournament_id, :user_id)
  end
end
