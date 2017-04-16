class AssociateUserWithPlayer < ActiveRecord::Migration
	def change
		remove_column :players, :name
		add_column :users, :name, :string
		add_reference :players, :user, index: true
	end
end
