import UIKit
import Goprojlib

class ViewController: UIViewController {
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        /* Calling the test function, off the main thread */
        DispatchQueue.global(qos: DispatchQoS.QoSClass.userInitiated).async {
            Goprojlib.GoprojlibRunChannels()
        }
        
        /* Animate the view to make sure the app didn't freeze anywhere */
        Timer.scheduledTimer(withTimeInterval: 1, repeats: true) { (_) in
            self.view.backgroundColor = UIColor.random
        }
    }
    
}

extension UIColor {
    static var random: UIColor {
        return UIColor(
            red: 0.5,
            green: .random(in: 0.5...0.8),
            blue: 1.0,
            alpha: 1.0
        )
    }
}
