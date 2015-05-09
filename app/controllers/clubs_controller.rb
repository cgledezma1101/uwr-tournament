class ClubsController < ApplicationController
  before_action :authenticate_user!
  load_and_authorize_resource

  # GET /clubs/:id
  def show
  end

  # POST /clubs
  def new
    @club = Club.new

    respond_to do |format|
      format.html{ render 'clubs/_new', layout: false }
    end
  end
end
