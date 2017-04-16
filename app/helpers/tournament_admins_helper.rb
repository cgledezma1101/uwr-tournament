module TournamentAdminsHelper
	# Allows the creation of a method that will render management actions over a tournament's admins
	#
	# @param [Tournament] tournament The tournament for which the actions will be performed
	#
	# @return [Proc] A method that can be called with a particular administrator and that will render the actions that can be performed on that administrator for the mentioned tournament
	def create_admin_actions_proc(tournament)
		Proc.new do |user|
			tournament_admin = TournamentAdmin.find_by(tournament: tournament, user: user)
			render 'tournaments/admin_actions', tournament_admin: tournament_admin
		end
	end
end
