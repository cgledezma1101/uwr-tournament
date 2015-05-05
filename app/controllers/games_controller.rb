class GamesController < ApplicationController
  before_action :authenticate_user!

  # POST /games
  #
  # Based on the parameters received, creates a new game and redirects the user
  # to the screen that visualizes it.
  def create
    @game = Game.new(new_game_params)
    @game.date = Time.now
    if @game.save
      redirect_to game_path(@game)
    else
      @teams = Team.all
      render 'games/new'
    end
  end

  # GET /games
  #
  # Renders the list of games played during the tournament
  def index
    @games = Game.all
  end

  # GET /
  # On the root of the application, allows the creation of a new game
  def new
    @teams = Team.all
    @game = Game.new
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
    def new_game_params
      params.require(:game).permit(:blue_team_id,
                                   :white_team_id,
                                   player_ids: [])
    end
end
