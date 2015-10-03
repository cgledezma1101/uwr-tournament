class Stage < ActiveRecord::Base
  belongs_to :tournament

  validates :tournament, presence: true
  validates :name, presence: true
end
