class GamesModelUpdate < ActiveRecord::Migration[6.0]
	def change
		remove_column :games, :date, :date
		add_column :games, :status, :string
	end
end
