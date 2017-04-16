class GameEvent < ActiveRecord::Base
	belongs_to :game

	validates :text, presence: true
end
