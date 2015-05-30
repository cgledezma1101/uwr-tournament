class PlayersController < ApplicationController
  before_action :authenticate_user!
  load_and_authorize_resource

  # GET /players/new?team_id=:team_id
  #
  # Renders a form that allows the addition of a player to a team
  #
  # @param [Integer] team_id Identifier of the team that will be updated
  def new
    @team = Team.find(params[:team_id])
    authorize! :update, @team

    current_players = @team.users
    clubs_members = @team.club_members
    @available_players = (clubs_members - current_players).sort{ |user0, user1| user0.name <=> user1.name }

    render 'players/_new.html.haml', layout: false
  end

  # POST /players
  #
  # Adds a player to a team
  def create
    if(@player.save)
      redirect_params = { notice: t('player.create_success') }
    else
      redirect_params = { alert: t('player.create_failure') }
    end

    redirect_to team_path(@player.team), redirect_params
  end

  private

  # Sanitazes arguments for the creation of  aplayer
  def create_params
    params.require(:player).permit(:team_id, :user_id)
  end
end
