class AssociateUserWithPlayer < ActiveRecord::Migration[6.0]
	def change
		remove_column :players, :name
		add_column :users, :name, :string
		add_reference :players, :user, index: true
	end
end
