from flask import Flask, request, jsonify

app = Flask(__name__)

foods = {
    "pizza": {
        "meal": "Pepperoni Pizza",
        "drink": "Cola"
    },
    "burger": {
        "meal": "Double Cheese Burger",
        "drink": "Milkshake"
    },
    "chicken": {
        "meal": "Grilled Chicken",
        "drink": "Orange Juice"
    },
    "rice": {
        "meal": "Jollof Rice with Chicken",
        "drink": "Chapman"
    }
}

@app.route("/recommend", methods=["POST"])
def recommend():
    data = request.get_json()

    food = data.get("food", "").lower()

    result = foods.get(
        food,
        {
            "meal": "Chef's Special",
            "drink": "Fresh Juice",
        },
    )

    return jsonify(result)

if __name__ == "__main__":
    app.run(debug=True, port=5000)