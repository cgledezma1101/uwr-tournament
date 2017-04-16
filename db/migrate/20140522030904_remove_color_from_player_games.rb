class RemoveColorFromPlayerGames < ActiveRecord::Migration
	def change
		remove_column :player_games, :color
	end
end
