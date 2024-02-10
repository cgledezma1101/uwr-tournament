class RemoveColorFromPlayerGames < ActiveRecord::Migration[6.0]
	def change
		remove_column :player_games, :color
	end
end
