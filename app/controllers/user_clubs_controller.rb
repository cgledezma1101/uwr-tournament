class UserClubsController < ApplicationController
  before_action :authenticate_user!
  load_and_authorize_resource

  # GET /user_clubs/:id/confirm_destroy
  #
  # Renders a confirmation message that allows the removal of the membership of a user to a club
  #
  # @param [Integer] id Identifier of the relationship that would be destroyed
  def confirm_destroy
    render 'user_clubs/_confirm_destroy', layout: false
  end

  # DELETE /user_clubs/:id
  #
  # Removes a user's membership from a club
  #
  # @param [Integer] id Identifier of the existing relationship
  def destroy
  end
end
