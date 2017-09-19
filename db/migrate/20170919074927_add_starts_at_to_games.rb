class AddStartsAtToGames < ActiveRecord::Migration
	def change
		add_column :games, :starts_at, :datetime
	end
end
