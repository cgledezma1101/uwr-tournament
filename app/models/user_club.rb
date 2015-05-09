class UserClub < ActiveRecord::Base
  belongs_to :user
  belongs_to :club

  validates :user, presence: true, uniqueness: { scope: :club }
  validates :club, presence: true
end
