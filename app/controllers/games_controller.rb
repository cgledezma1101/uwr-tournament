class GamesController < ApplicationController
  # POST /games
  #
  # Based on the parameters received, creates a new game and redirects the user
  # to the screen that visualizes it.
  def create
    binding.pry
  end

  # GET /
  # On the root of the application, allows the creation of a new game
  def new
    @teams = Team.all
    @game = Game.new
  end
end
