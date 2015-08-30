class TournamentsController < ApplicationController
  before_action :authenticate_user!
  load_and_authorize_resource

  # GET /tournaments/new
  #
  # Provides a form that contains the details to create a new tournament
  def new
    render 'tournaments/_new', layout: false
  end

  # POST /tournaments
  #
  # Allows the creatpion of a new tournament
  def create
    if(@tournament.save)
      TournamentAdmin.create(user: current_user, tournament: @tournament)

      redirect_to tournament_path(@tournament)
    else
      redirect_to root_path, alert: t('tournament.create_error')
    end
  end

  private

  # Sanitizes arguments for mass initialization
  def create_params
    params.require(:tournament).permit(:name, :start_date, :end_date)
  end
end
