class AddIsActiveToPlayers < ActiveRecord::Migration[6.0]
	def change
		add_column :players, :is_active, :boolean
	end
end
