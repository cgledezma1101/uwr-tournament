class AddWinningColorToGames < ActiveRecord::Migration[6.0]
	def change
		add_column :games, :winning_color, :string
	end
end
