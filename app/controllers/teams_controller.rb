class TeamsController < ApplicationController
  before_action :authenticate_user!

  # GET /teams/:id/players
  #
  # For the specified team, searches it's players and renders a form associated
  # to them. This action answers esclusively to JS calls.
  #
  # @param [Integer] id Identifier of the team whos players are desired
  def players
    @team = Team.find(params[:id])
    @players = @team.players
    @game = Game.new
  end

  # GET /teams/new
  #
  # Renders the form that will allow the creation of a new team for a club
  #
  # @param [Integer] club_id Identifier of the club to which the team will be added
  def new
    @club = Club.find(params[:club_id])
    authorize! :edit, @club

    @team = Team.new
    respond_to do |format|
      format.html{ render 'teams/_new', layout: false }
    end
  end

  # POST /teams
  #
  # Allows for the creation of a new team, associating it to a club
  def create
    club = Club.find(params[:team][:club_id])
    authorize! :edit, club

    team = Team.create(create_params)

    if(team.valid?)
      redirect_to team_path(team)
    else
      redirect_to root_path, alert: t('team.standard_save_error')
    end
  end

  # GET /teams/:id
  #
  # Displays the general information of a particular team
  def show
    @team = Team.find(params[:id])
    authorize! :show, @team
    @club = @team.club
  end

  private

  # Parameter sanitazer for the create method
  def create_params
    params.require(:team).permit(:name, :club_id)
  end
end
