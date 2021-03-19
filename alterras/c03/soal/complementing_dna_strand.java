import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.regex.*;



class Result {

    /*
     * Complete the 'dnaComplement' function below.
     *
     * The function is expected to return a STRING.
     * The function accepts STRING s as parameter.
     */
public static String makeComplement(String dna) {
    String output = "";
    String [] dnaArray = dna.split("");
    for(String base : dnaArray){
        output += getDNACompliment(base);
        }
    return output;
  }
    public static String getDNACompliment(String base) {
    // Write your code here
String complement = "";
      switch (base){
        case "A": complement = "T";
                  break;
        case "T": complement = "A";
                 break;
        case "G": complement = "C";
                  break;
        case "C": complement = "G";
                  break;
      }
      return complement;
    }

}

public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s = bufferedReader.readLine();

        String result = Result.makeComplement(s);

        bufferedWriter.write(result);
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
