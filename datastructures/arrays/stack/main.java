import java.util.Stack;

class Main {
	private static int calPoints(String[] operations) {
		int sum = 0;
		Stack<Integer> scoreRecord = new Stack<>();

		for (String val : operations) {
			switch (val) {
				case "+":
					int firstElem = scoreRecord.peek();
					int secondElem = scoreRecord.get(scoreRecord.size() - 2);
					scoreRecord.push(firstElem + secondElem);
					break;
				case "D":
					scoreRecord.push(scoreRecord.peek() * 2);
					break;
				case "C":
					if (!scoreRecord.isEmpty()) {
						scoreRecord.pop();
					}
					break;
				default:
					int elem = Integer.parseInt(val);
					scoreRecord.push(elem);
					break;

			}

		}

		for (int val : scoreRecord) {
			sum += val;
		}

		return sum;
	}

	public static void main(String[] args) {
		String[] ops = new String[] { "1", "2", "+", "C", "5", "D" };
		System.out.println(calPoints(ops));

	}
}
