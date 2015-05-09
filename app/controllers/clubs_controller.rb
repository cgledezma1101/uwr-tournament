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
    @club.admins << current_user

    binding.pry

    if @club.valid?
      @club.save
      redirect_to clubs_path(@club)
    else
      redirect_to root_path, alert: t('club.standard_save_error')
    end
  end

  private

  def create_params
    params.require(:club).permit(:name)
  end
end
