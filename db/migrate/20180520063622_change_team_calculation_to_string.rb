class ChangeTeamCalculationToString < ActiveRecord::Migration[6.0]
  def change
    remove_column :games, :blue_team_calculation, :integer
    remove_column :games, :white_team_calculation, :integer

    add_column :games, :blue_team_calculation, :string
    add_column :games, :white_team_calculation, :string
  end
end
