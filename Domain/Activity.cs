namespace Domain
{
    public class Activity //Table
    {
        //Guid (Globally Unique Identifier) é um tipo fornecido pelo .NET para representar um identificador global único. É frequentemente usado como identificador único para objetos ou entidades.
        public Guid Id { get; set; } //get and set the value methods

         public string Title { get; set; } //collumn

         public DateTime Date { get; set; }  //collumn

         public int Description { get; set; }  //collumn

         public int Category { get; set; }  //collumn

         public string City {get; set;}  //collumn

         public int Venue { get; set; }  //collumn
    }
}