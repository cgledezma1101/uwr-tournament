class GameEventsController < ApplicationController
  load_and_authorize_resource
  before_action :authenticate_user!

  # POST /game_events
  #
  # Allows the creation of an event within a game
  #
  # @param [Integer] game_id The identifier of the game to which the event will be associated to
  # @param [String] text The contents of the event being created
  def create
    @game_event.save
  end

  # DELETE /game_events/:id
  #
  # Allows the deletion of an event in a game
  #
  # @param [Integer] id Identifier of the feed to be destroyed
  def destroy
    @game_event.destroy
  end

  private

  # Sanitizes the parameters to allow mass assignment in the create action
  def create_params
    params.require(:game_event).permit(:game_id, :text)
  end
end
