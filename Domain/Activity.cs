namespace Domain
{
    public class Activity
    {
        //Guid (Globally Unique Identifier) é um tipo fornecido pelo .NET para representar um identificador global único. É frequentemente usado como identificador único para objetos ou entidades.
        public Guid Id { get; set; } //get and set the value methods

         public string Title { get; set; }

         public DateTime Date { get; set; }

         public int Description { get; set; }

         public int Category { get; set; }

         public string City {get; set;}

         public int Venue { get; set; }
    }
}