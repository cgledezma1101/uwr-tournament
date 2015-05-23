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

  # GET /teams/new
  #
  # Renders the form that will allow the creation of a new team for a club
  #
  # @param [Integer] club_id Identifier of the club to which the team will be added
  def new
    @club = Club.find(params[:club_id])
    authorize! :edit, @club

    render 'teams/_new', layout: false
  end

  # POST /teams
  #
  # Allows for the creation of a new team, associating it to a club
  def create
    club = Club.find(params[:team][:club_id])
    authorize! :edit, club

    if(@team.save)
      redirect_to team_path(@team)
    else
      redirect_to club_path(club), alert: t('team.standard_save_error')
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

  private

  # Parameter sanitazer for the create method
  def create_params
    params.require(:team).permit(:name, :club_id)
  end
end
