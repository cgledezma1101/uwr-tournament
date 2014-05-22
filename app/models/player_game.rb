class PlayerGame < ActiveRecord::Base
  belongs_to :player
  belongs_to :game

  validates :player, presence: true, uniqueness: { scope: :game }
  validates :game, presence: true
end
