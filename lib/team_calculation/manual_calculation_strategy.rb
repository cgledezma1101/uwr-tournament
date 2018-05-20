class TeamCalculation::ManualCalculationStrategy < TeamCalculation::CalculationStrategy
	def calculate_team(game, team_color)
		if (team_color == PlayerGame::WHITE_TEAM)
			return game.white_team_id.nil? ? nil : Team.find(game.white_team_id)
		elsif (team_color == PlayerGame::BLUE_TEAM)
			return game.blue_team_id.nil? ? nil : Team.find(game.blue_team_id)
		end

		raise "Invalid team color provided to calculation strategy: #{team_color}"
	end
end