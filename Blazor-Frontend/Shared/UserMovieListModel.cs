using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

using System.Globalization;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;

namespace Blazor_Frontend.Shared
{
    public partial class UserMovieListModel
    {
        [JsonProperty("Name")]
        public string Name { get; set; }

        [JsonProperty("Movies")]
        public Movie[] Movies { get; set; }

        public override string ToString()
        {
            var Titles = "Movie List: ";
            foreach (Movie movie in Movies) {
                Titles += movie;
            }
            return "Name: " + Name + Titles;
        }
    }

    public partial class Movie
    {
        [JsonProperty("Name")]
        public string Name { get; set; }

        [JsonProperty("Raiting")]
        public int Raiting { get; set; }

        [JsonProperty("ID")]
        public string ID { get; set; }

        [JsonProperty("Crew")]
        public Person[]? Crew { get; set; }
        public override string ToString()
        {
            return Name + ", ";
        }
    }

    public partial class Person : IEquatable<Person>
    {
        public string Name { get; set; }
        public string Roll { get; set; }

        public Person(string Name, string Roll)
        {
            this.Name = Name;
            this.Roll = Roll;
        }

        public bool Equals(Person other)
        {
            if (this.Name == other.Name && this.Roll == other.Roll)
            {
                return true;
            } else
            {
                return false;

            }
        }

        public override bool Equals(object obj)
        {
            if (obj == null) {
                return false;
            }
            Person personObj = obj as Person;
            if (personObj == null)
            {
                return false;
            } else
            {
                return Equals(obj as Person);
            }
        }

        public override int GetHashCode()
        {
            return (this.Name.GetHashCode() + this.Roll.GetHashCode()).GetHashCode();
        }
    }
}

