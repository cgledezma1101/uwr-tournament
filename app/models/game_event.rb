class GameEvent < ApplicationRecord
	belongs_to :game

	validates :text, presence: true
end
