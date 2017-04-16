class GamesModelUpdate < ActiveRecord::Migration
	def change
		remove_column :games, :date, :date
		add_column :games, :status, :string
	end
end
