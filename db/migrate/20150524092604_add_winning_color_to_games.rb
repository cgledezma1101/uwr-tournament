class AddWinningColorToGames < ActiveRecord::Migration
	def change
		add_column :games, :winning_color, :string
	end
end
