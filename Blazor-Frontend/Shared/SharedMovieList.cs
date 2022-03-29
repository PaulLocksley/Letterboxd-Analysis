using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Blazor_Frontend.Shared
{

    public struct SharedMovieObject
    {
        //I feel like theres a better way to represent this relationship using the data I allready have.
        public UserRaiting[] UserRaitings;
        public Movie Movie;
        public string ID;
        public int AverageRaitingDifference;
        public SharedMovieObject(UserRaiting[] UserRaitings, Movie Movie, string ID)
        {
            this.ID = ID;
            this.UserRaitings = UserRaitings;
            this.Movie = Movie;
            var userRaitingTotal = 0;
            for (var i = 1; i < UserRaitings.Length; i++)
            {
                userRaitingTotal += UserRaitings[0].Raiting - UserRaitings[i].Raiting;
            }
            //this.AverageRaitingDifference = (int)Math.Round((float)Math.Abs(userRaitingTotal) / (float)(UserRaitings.Length-1), MidpointRounding.AwayFromZero);
            this.AverageRaitingDifference = (int)Math.Round((float)userRaitingTotal / (float)(UserRaitings.Length - 1), MidpointRounding.AwayFromZero);
        }
    }
    public struct UserRaiting
    {
        public string UserName;
        public int Raiting;
        public UserRaiting(string userName, int Raiting)
        {
            this.UserName = userName;
            this.Raiting = Raiting;
        }
    }

    public enum LoadingStates
    {
        fetching,
        parsing,
        done
    }
}