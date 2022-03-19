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

    public partial class Person
    {
        public string Name { get; set; }
        public string Roll { get; set; }
    }
}

