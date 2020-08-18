package com.udacity.sandwichclub.utils;

import android.util.Log;

import com.udacity.sandwichclub.model.Sandwich;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.util.ArrayList;

public class JsonUtils {
    public static Sandwich parseSandwichJson(String json) {
        Sandwich sandwich = null;
        try {
            JSONObject jsonObject = new JSONObject(json);

            JSONObject nameObj = jsonObject.getJSONObject("name");
            String mainName = nameObj.getString("mainName");
            JSONArray alsoKnowAsArray = nameObj.getJSONArray("alsoKnownAs");

            String placeOfOrigin = jsonObject.getString("placeOfOrigin");
            String description = jsonObject.getString("description");
            String image = jsonObject.getString("image");
            JSONArray ingredientsArray = jsonObject.getJSONArray("ingredients");

            ArrayList<String> alsoKnownAsList = jsonArrayToArrayList(alsoKnowAsArray);
            ArrayList<String> ingredientsList = jsonArrayToArrayList(ingredientsArray);

            sandwich = new Sandwich(mainName, alsoKnownAsList, placeOfOrigin, description, image, ingredientsList);

        } catch (JSONException e) {
            e.printStackTrace();
        }
        return sandwich;
    }

    private static ArrayList<String> jsonArrayToArrayList(JSONArray values) {
        ArrayList<String> results = new ArrayList<String>();
        try {
            for(int index = 0; index < values.length(); index++) {
                results.add((String) values.get(index));
            }
        } catch(JSONException e) {
            e.printStackTrace();
        }
        return results;
    }
}
