class User < ActiveRecord::Base
  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :trackable, :validatable

   has_many :players

   has_many :user_clubs, dependent: :destroy
   has_many :clubs, through: :user_clubs

   has_many :club_admins
   has_many :administrated_clubs, through: :club_admins
end
