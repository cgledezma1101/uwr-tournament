class TeamsController < ApplicationController
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
end
