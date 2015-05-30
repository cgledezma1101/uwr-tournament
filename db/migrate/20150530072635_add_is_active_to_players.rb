class AddIsActiveToPlayers < ActiveRecord::Migration
  def change
    add_column :players, :is_active, :boolean
  end
end
