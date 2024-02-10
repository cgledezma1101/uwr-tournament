class AddStageIdToGame < ActiveRecord::Migration[6.0]
	def change
		add_column :games, :stage_id, :integer
	end
end
