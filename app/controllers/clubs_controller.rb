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

  # GET /clubs/:id/new_admin
  #
  # Presents an interface to add a new administrator to the club
  def new_admin
    respond_to do |format|
      format.html{ render 'clubs/_new_admin', layout: false }
    end
  end

  # POST /clubs/:id/create_admin
  #
  # Adds a new administrator to the club
  #
  # @param [Integer] id Identifier of the club to be updated
  # @param [Integer] user[id] Identifier of the user that's going to be added as administrator
  def create_admin
    user = User.find(params[:user][:id])
    membership = UserClub.find_by(user: user, club: @club)

    if membership.nil?
      redirect_params = { alert: t('club.errors.admin_must_be_member') }
    else
      admin = ClubAdmin.new(club: @club, user: user)

      if admin.valid?
        membership.destroy
        admin.save

        redirect_params = { notice: t('club.admin_create_success') }
      else
        redirect_params = { alert: t('club.errors.admin_not_added') }
      end
    end

    redirect_to club_path(@club), redirect_params
  end

  # GET /clubs/:id/edit
  #
  # Returns the form required to update the details of the club
  #
  # @param [Integer] id Identifier of the club that will be updated
  def edit
    respond_to do |format|
      format.html{ render 'clubs/_new', layout: false }
    end
  end

  # POST/PUT /clubs/:id
  #
  # @param [Integer] id Identifier of the club to be edited
  # @param [String] club[name] New name for the club
  def update
    if @club.update_attributes(update_params)
      redirect_params = { notice: t('club.update_successful') }
    else
      redirect_params = { alert: t('club.update_failure') }
    end

    redirect_to club_path(@club), redirect_params
  end

  private

  def create_params
    params.require(:club).permit(:name)
  end

  def update_params
    params.require(:club).permit(:name)
  end
end
