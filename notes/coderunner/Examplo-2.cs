 // -- CODE RUNNER BELOW --

//https://onecompiler.com/csharp/

using System;

namespace ArrayCollection
{
    public class CollectionHandler
    {
        // A expressão new int[] é necessária para inicializar um array em C#.
        public int[] Numbers { get; set; } = new int[] { 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 };
        public string[] Strings { get; set; } = new string[] { "1", "2", "3", "4", "5", "6", "7", "8", "9", "10" };

        public void SumNumberCollection()
        {
            int sum = 0;
            foreach (var item in Numbers)
            {
                sum += item;
            }
            Console.WriteLine($"Sum of Numbers collection: {sum}");
        }

        public void SumStringCollection()
        {
            int sum = 0;
            foreach (var item in Strings)
            {
                sum += int.Parse(item);
            }

            Console.WriteLine($"Sum of Strings collection: {sum}");
        }

        public void SummAllCollection()
        {
            int totalSum = 0;

            // Somando os valores de Numbers
            foreach (var number in Numbers)
            {
                totalSum += number;
            }

            // Somando os valores de Strings (convertendo para inteiro)
            foreach (var str in Strings)
            {
                totalSum += int.Parse(str);
            }

            Console.WriteLine($"Sum of all collections: {totalSum}");
        }
    }

    class Program
    {
        static void Main()
        {
            var collectionHandler = new CollectionHandler();
            collectionHandler.SumNumberCollection();
            collectionHandler.SumStringCollection();
            collectionHandler.SummAllCollection(); // Chama o método que soma os dois arrays
        }
    }
}
