class GamesController < ApplicationController
  load_and_authorize_resource
  before_action :authenticate_user!

  # POST /games
  #
  # Based on the parameters received, creates a new game and redirects the user
  # to the screen that visualizes it.
  def create
    @game.status = Game::STATUS_READY
    if @game.save
      redirect_params = { notice: t('game.create_success') }
    else
      redirect_params = { alert: t('game.create_fail') }
    end

    redirect_to stage_path(@game.stage), redirect_params
  end

  # DELETE /games/:id
  #
  # Destroys the specified instance of a game
  #
  # @param [Integer] id Identifier of the game to be destroyed
  def destroy
    @game.destroy
    redirect_to stage_path(@game.stage), notice: t('game.destroy_success')
  end

  # GET /games/new?stage_id=:stage_id
  #
  # Renders a form that would allow the creation of a game in a stage
  #
  # @param [Integer] stage_id Identifier of the stage to which this game will be added
  def new
    stage = Stage.find(params[:stage_id])
    authorize! :update, stage

    @game.stage = stage
    @teams = stage.tournament.teams.order(:name)

    render 'games/_new', layout: false
  end

  # GET /games/:id
  #
  # Displays a game and allows basic mannipulation over it
  #
  # @param [Integer] id The identifier of the game to be displayed
  def show
    @game = Game.find(params[:id])
    @blue_players = @game.blue_players
    @white_players = @game.white_players
    @score = Score.new
  end

  private
    # This defines the attributes that are permitted when creating a new game
    def create_params
      params.require(:game).permit(:blue_team_id, :white_team_id, :stage_id)
    end
end
