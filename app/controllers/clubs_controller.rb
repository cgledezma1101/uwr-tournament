class ClubsController < ApplicationController
  before_action :authenticate_user!
  load_and_authorize_resource

  # GET /clubs/:id
  def show
  end

  # GET /clubs/new
  def new
    respond_to do |format|
      format.html{ render 'clubs/_new', layout: false }
    end
  end

  # POST /clubs
  def create
    if @club.valid?
      @club.save
      ClubAdmin.create(club: @club, user: current_user)

      redirect_to club_path(@club)
    else
      redirect_to root_path, alert: t('club.standard_save_error')
    end
  end

  # GET clubs/:id/new_admin
  #
  # Presents an interface to add a new administrator to the club
  def new_admin
    respond_to do |format|
      format.html{ render 'clubs/_new_admin', layout: false }
    end
  end

  private

  def create_params
    params.require(:club).permit(:name)
  end
end
