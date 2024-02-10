class AddStartsAtToGames < ActiveRecord::Migration[6.0]
	def change
		add_column :games, :starts_at, :datetime
	end
end
