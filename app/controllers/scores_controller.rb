class ScoresController < ApplicationController
  before_action :authenticate_user!
  
  # POST /scores
  #
  # Method called when a player scores a goal on a game
  #
  # @param [Integer] game_id Identifier of the game where the score was made
  # @param [Integer] player_id Identifier of the player that scored
  def create
    @game = Game.find params[:score][:game_id]
    @player = Player.find params[:score][:player_id]

    @score = Score.create(game: @game, player: @player)
  end
end
