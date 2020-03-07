class TeamCalculation::CalculationStrategy
	# Determines the team to be used for the specified color
	#
	# @param [Game] game The game for which the  team should be calculated
	# @param [String] team_color PlayerGame::BLUE_TEAM or PlayerGame::WHITE_TEAM, depending on which team wants to be determined
	# @return Team The calculated team
	def calculate_team(game, team_color)
		raise "calculate_team not implemented"
	end
end