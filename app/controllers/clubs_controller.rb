class ClubsController < ApplicationController
	before_action :authenticate_user!
	load_and_authorize_resource

	# GET /clubs/:id
	def show
	end

	# GET /clubs/new
	def new
		render 'clubs/_new', layout: false
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

	# GET /clubs/:id/edit
	#
	# Returns the form required to update the details of the club
	#
	# @param [Integer] id Identifier of the club that will be updated
	def edit
		render 'clubs/_new', layout: false
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
