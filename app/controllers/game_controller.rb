class GameController < ApplicationController
  # /
  # On the root of the application, allows the creation of a new game
  def new
    @teams = Team.all
    @game = Game.new
  end
end
