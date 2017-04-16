class AddTeamColorToPlayerGame < ActiveRecord::Migration
	def change
		add_column :player_games, :team_color, :string
	end
end
