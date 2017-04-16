class StagesController < ApplicationController
	before_action :authenticate_user!
	load_and_authorize_resource

	# POST /stages
	#
	# Allows the creation of a new stage
	#
	# @param [Integer] tournament_id Identifier of the tournament the stage will belong to
	# @param [Integer] name Name the stage is going to have
	def create
		if @stage.save
			redirect_params = { notice: t('stage.save_success') }
		else
			redirect_params = { alert: t('stage.save_fail') }
		end

		redirect_to tournament_path(@stage.tournament), redirect_params
	end

	# DELETE /stage/:id
	#
	# Destroys the specified instance of the stage
	#
	# @param [Integer] id Identfier of the stage to be destroyed
	def destroy
		@stage.destroy
		redirect_to tournament_path(@stage.tournament), notice: t('stage.destroy_successful')
	end

	# GET /stages/new?tournament_id=:tournament_id
	#
	# Renders a form that would allow the creation of a stage inside a tournament
	#
	# @param [Integer] tournament_id Identifier of the tournament this stage will be added to
	def new
		tournament = Tournament.find(params[:tournament_id])
		authorize! :update, tournament

		@stage.tournament = tournament
		render 'stages/_new', layout: false
	end

	private

	# Returns the sanitized set of parameters for mass assignment of a new stage
	def create_params
		params.require(:stage).permit(:tournament_id, :name)
	end
end
