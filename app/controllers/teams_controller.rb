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
end
